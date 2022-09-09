## GoTest-2: Documentation

# Functions:

benchmarkParallelsum: 
  Takes a slice created from user input, and a channel. The function will sum values in the slice and return result to a channel
  
benchmarkSlicestable:
  Takes an integer value from user, generates a slice of random x numbers, and analyzes performance of sort.Slice and sort.SliceStable
  
main:
  Driver that takes user input from command line and sets up the benchmarking for the parallel sum function. To do this the function creates a nums array from the x user value, and calls the parallel sum function with different slices and channels. The main function also calls the slicestable function which has the benchmarking built in.
  
  
# How to use ( how it is supposed to be used T-T):
  Simply clone the repo on your local machine, and move into the project directory.
  Here type "go run main.go" in your cmd line to start the program.
  Enter an integer value of your choice and observe the benchmarking results.
  
# Reflection
  The program ran into some issues when implementing the parallel functionality, although in theory I thought my implementation would work--it seems I still need some practice/understanding of go routines, channels, and concurrency in general to make this work. However, Im still satisfied with what I did in a short time and will work to find the solution in my own time.
  
#Thanks for reading
