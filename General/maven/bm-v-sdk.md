# what is the difference between a bom and an sdk

## bom

[Maven Docs on it](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html)

Control how the dependencies of a multi-module project are versioned in user projects.
Controlling them keeps them in sync.
Keeping them in sync reduces the possibility of strange errors due to older versions being used or runtime classloader errors. (NoClassDefFoundError, NoSuchMethodError)
Bom is specified in dependencyManagement section of root pom.

```xml
<depenendencyManagement>
	<dependency>
		<groupid>...</groupid>
		<artifactId>...</artifactId>
		<version>...</version>
		<type>pom</type>
		<scope>import</scope>
	</dependency>
</depenendencyManagement>
```
<scope> element in dependency:

import: instead of importing me import my dependencies.
test: this dependency is only available at the test phase (classpath)

<type> element in dependency:

refers to packaging type of artifact, defaults to jar.

... 

Dependencies have to be individually added to the project, without their versions

```xml
<dependency>
	<groupId>...</groupId>
	<artifactId>...</artifactId>
</dependency>
```

### Questions:

1. type element in pom?

## sdk

grouping of dependencies which the versions of are controlled and imported for the user.

### Definition from Google

"A software development kit (SDK or "devkit") is typically a set of software development tools that allows the creation of applications for a certain software package, software framework, hardware platform, computer system, video game console, operating system, or similar development platform."

### sdk in maven terms

a set of software packages that allow the creation of applications

### taf-sdk

a set of taf modules which allow the creation of automated tests.

## Thoughts

the sdk makes it easier for users, but bloats the classpath?
the bom will minimize the class path but the users must specify the individual dependencies.