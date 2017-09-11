Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-03-13T11:33:29+00:00

###### Pattern API ######
Created Sunday 13 March 2016

Sequence of usage:
	Compile a pattern, a regular expression
	create a matcher from the pattern passing the string to match.
	use matcher to verify comparison

Pattern contains a convenience matches method for when the pattern will only be matched once

Instances of Pattern are thread safe, instances of Matcher are not.

https://docs.oracle.com/javase/7/docs/api/java/util/regex/Pattern.html, useful for summary of regular expression syntax

Fun way to learn regex: https://regexcrossword.com/

http://tutorials.jenkov.com/java-regex/pattern.html

Text can be split on matched pattern using pattern.split
	pattern matched is not included in resulting array

Pattern pattern # Pattern.compile(<pattern to match>);
Matcher matcher # pattern.matcher(<string to try and match pattern in>);
matcher.matches(); boolean whether there was a match or not.

above 3 lines can be condensed into 
Pattern.matches(<pattern>,<text>); returns boolean

memory over head if same pattern is compiled repeatedly
Compile the pattern once and reuse it.

String[] split # pattern.split(<text>);

String patternString # pattern.pattern();

Matcher methods
	lookingAt
		Try to match with the start of the text.
	find
		find the first occurance of the regex in the text.
		every subsequent call finds the next occurance.
	start
		starting index of the occurance of the regex in the string
		just like find they return the start index of the subsequent find for the next method call
	end
		index just after the last character of the occurance of the regex in the string
		just like find they return the end index of the subsequent find for the next method call
	reset
		reset the matcher internally
		reset is overloaded to take a string which is the new text to match against
	**group!!!!!**
		extract the matching strings
		give an index of the one you want to extract
			0 is the whole regex
			so use 1
		there can be multiple groups in a regex
			so indexes are for each match
		keep calling find to get subsequent matches
		you can have groups inside groups
			1 is the big group
			2 is the first inner group
			...
		( ) marks a group in the regex
	replaceAll
		replace all occurances of the matched regex
		matcher is reset before replace being called.
	replaceFirst
		replace the first occurance of the matched regex
	appendReplacement
		append text + occurance to stringbuffer and replace occurance
	appendTail
		append everything after the last occurance to the stringbuffer

http://tutorials.jenkov.com/java-regex/syntax.html
