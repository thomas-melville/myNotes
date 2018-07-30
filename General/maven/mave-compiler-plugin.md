# maven compiler plugin

## configure to fail on compiler warnings

```xml

<plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.6.0</version>
                <configuration>
                    <source>1.8</source>
                    <target>1.8</target>
                    <failOnWarning>true</failOnWarning>
                    <compilerArgument>-Xlint</compilerArgument>
                </configuration>
            </plugin>

```