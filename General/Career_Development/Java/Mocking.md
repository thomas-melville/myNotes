# Mocking

Classes don't exist in isolation, they use other classes.
Those other classes may talk to external systems which won't work in a unit test environment
Because of this we need to mock the collaborators of a class

Mockito is the framework of choice

PowerMockito is used when mocking of static & private methods is involved
	uses bytecode manipulation
	comes with it's own runner

1. @RunWith(PowerMockRunner.class)
2. @PrepareForTest(< classes with static methods to be mocked>)
3. mockStatic(<individual classes>)
