# Spring Util classes

Assert
  bunch of assertions methods for different types
  throws IllegalArgumentException if the assertion fails.
  generally for internal use
BeanUtils
  get property descriptors of a class
  create an instance in the ApplicationContext.
PropertyDescriptor
  about a class
ClassUtils
  methods which use reflection to get runtime information about the instance
SystemPropertyUtils
  resolvePlaceHolders, use ${} in strings
FileCopyUtils
  another class of utilities for copying files
  copy file lines into a string
ServletRquestUtils
  get parameters from a request
  when not using mapping to an object
WebUtils
  Get session information
  Get temporary directory information in the context of the request
RequestContextUtils
  find applicationcontext, only needed in very old spring code
