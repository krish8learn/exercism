package erratum

func Use(opener ResourceOpener, input string) (frobErr error) {

	// opening the resource by calling function --> type ResourceOpener func() (Resource, error)
	resource, resourceError := opener()

	//handle the error may be produced here
	if resourceError != nil {
		//it may fail with an error of type TransientError
		for isTransientError(resourceError) {
			//transient error occurred
			//open again
			resource, resourceError = opener()
			//if the error is different than , return it
			if !isTransientError(resourceError) && resourceError != nil {
				//not transient error
				return resourceError
			}
			// again transient error occured, check again hope transient does not occurr
		}
	}

	//whatever happen resource must be closed, at end of the day
	defer resource.Close()

	//this function objctive is to handle panic occurred from Frob()
	// var frobErr error
	defer func() {
		//recover from panic
		recovered := recover()
		switch recovered.(type) {
		case nil:
			// no panic
		case FrobError:
			//deFrob it
			resource.Defrob(recovered.(FrobError).defrobTag)
			frobErr = recovered.(FrobError).inner
		default:
			//different error
			frobErr = recovered.(error)
		}

	}()

	//call the Frob(), which can produce panic
	resource.Frob(input)
	return frobErr
}

func isTransientError(err error) bool {
	//as we are re-opening the resource
	if err == nil {
		return false
	}
	//using type assertion to check error type
	switch err.(type) {
	case TransientError:
		return true
	default:
		return false
	}
}

//this is the most hardest problem I solved in exercism, maybe I have not understand this fully, working on it for 4 days,  whatever!! it works, leaving it and moving on to the next one, need new problems to solve 
