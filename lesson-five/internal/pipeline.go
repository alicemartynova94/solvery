package internal

type (
	In  = <-chan interface{}
	Out = In
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = stage(stageWrapper(in, done))
	}

	return stageWrapper(in, done)
}

func stageWrapper(in In, done In) Out {

	result := make(chan interface{})

	go func() {
		defer close(result)

		for {
			select {
			case <-done:
				return

			case v, ok := <-in:
				if !ok {
					return
				}

				select {
				case result <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return result
}
