package erratum

func Use(opener func() (Resource, error), input string) (err error) {
	resource, err := opener()

	// Use should keep trying if a transient error is returned on open.
	for err != nil {
		_, okTrans := err.(TransientError)
		// _, okFrob := err.(FrobError)
		// Use should fail if a non-transient error is returned on open.
		if !okTrans {
			return err
		}
		resource, err = opener()
	}
	defer resource.Close()
	// Use should call Defrob and Close on FrobError panic from Frob
	// and return the error.
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()
	// Use should not return an error on the "happy" path.
	resource.Frob(input)
	return nil
}
