# integration_testing
This project is very small and therefore might not be the perfect example for the assignment. Integration testing should be about testing 
different modules that all have different purposes. However we are still able to reflect over which approach and method should be used for minor systems as well as reflect over what we would have done in a much larger project with actual modules. 

## Testing plan
In integration testing we are really looking to make sure that all our modules work correctly when put together. That means that each module might work as expected in isolation but when combined with other modules we get defects and bugs. This is very likely to occur because in larger systems the different modules are coded by different programmers that all have a specific way of doing things. It is therefore crucial to test whether the modules work in unity.

In this example code, we are going with the big bang approach because we are working with a very small example. 

The big bang approach has a big advantage in small systems. All components are integrated together at once and then tested. We could just imagine that all the functions in *postgres.go* are individual modules that all take care of getting different kind of information from the database. Of course, this is not the actual case but let’s just imagine this is the case. 
We then test all these “modules” in our test file and make sure they all collect the correct data from the database. Again, it would be more realistic if the “modules” spoke to each other and was dependent on the data they got from each other but let’s just imagine that they are. 
We have now tested the example in a big bang approach.

The disadvantages to this approach would be the difficulty in localizing faults or bugs. Putting everything together at once would also mean you could easily miss interfaces to be tested. 
With this approach, it would also mean the testing could only begin after all the modules are designed and the testing team would have less time to test the solution. 
It makes sense to use the big bang approach in this case but in any larger system it would be preferred to use the bottom up or top down integration strategy. 

You could also argue that this example is the start of a bottom up approach. The modules that talk to the database was tested first and if we were to add a frontend we would then test this part later after the bottom part of the application were created and tested first.

The integration method we use are automated tests. All tests are automated in the code and can all be run together whenever we want to. 
That means we are using white-box-testing in this example. 
We could also choose to suppliment with black-box-testing and focus on whether the application accomplished the tasks it was created to solve.
If we had a GUI, a webpage for example, we could test whether the modules provided the information we were expecting to see in the application frontend.
That would mean we were using acceptance testing. 

## Test cases

Test Case ID | Test case description | Prerequisites | Test steps | Test data | Expected Result | Actual Result | Status | 
---|---|---|---|---|---|---|---|
TC001 | The objective of this test is to verify that a triangle with **three** sides of equal length will return 'equilateral triangle' | The user should obtain three equal integers | Enter the three integers to the program and run it | The input should all be integers. Any other datatypes are not valid | 1. The program should return 'equilateral triangle' if all three integers are equal to oneanother <br><br> 2. The program should return 'invalid input' if the input was not all integers | 1. The input was valid and the program returned 'equilateral triangle' <br><br> 2. If the input was invalid the program will return 'invalid input' | Success |


## Test results
![](https://github.com/xNoga/integration_testing/blob/master/images/Screen%20Shot%202018-04-16%20at%2014.16.15.png)
