# Prompt Engineering

What you get out of GenAI all depends on what you ask it and how!
Prompt Engineering is a skill set in it's own right!

## zero, single, few-shot learning approach

zero - ask the model to perform a certain task and expecting the model to understand how it should answer and what is being asked.
  prompt wont contain examples or demonstrations.

one/few-shot - give some examples of the desired behaviour to it and only then asking it do to a task which is closely related to those examples.
  provide examples/demonstrations in the prompt
  The model takes the example and uses that to refine its output
  You can provide one or as many examples as you like

chain-of-thought prompting
  in your examples provide intermediate reasoning steps, not just the answer.
  Combine it with few-shot to get better results.

zero-shot chain of thought
  Add "let's think step by step" to the end of your prompt

## Meta prompting

Advanced prompting technique, focuses on the structural and syntactical aspects of tasks and problems

## Input parameters

### Temperature

This term defines the level of exploration by the LLM.
0 - when you want reliability and predictability
0.x - when you want variety
I see it's defined in the REST API

### Role

You can structure your prompt using different roles: system, user and assistant.

system gives the LLM some context
user is the prompt from the user
assistant is the model response

## 7 building blocks of Prompt Engineering

1. Target Orientation
    prompt should have a clear goal
    specific but not too specific.
    Analyse X and create Y for Z
2. Clarity and precision
    we need to tell the AI what we want and how we want it
3. Context
    prompt should provide necessary context.
    Who is the audience, from what point of view, ...
4. Length
    Not too short, not too long. :-)
5. Language and style
    Should match the wanted language and style of the outcome
    State what language and style you want the response in.
    humorous, sarcastic, serious, ...
6. Adaptability
    responses can be refined by adjusting following prompt
7. Use of keywords
    can specify wanted aspects of the response.
    style, topics, ...

## Evaluating the responses

### RACCCA framework

Relevance
Accuracy
Completeness
Clarity
Coherence
Appropriateness


## Refining the response

First response is rarely exactly what you want.
You can ask follow up prompts to refine the response.

1. Analysis prompts
    Ask for more details on specific aspects of the response
2. Feedback prompts
    Ask for more answers on specific aspects of the response
3. Expansion prompts
    Expand on specific parts of the response
4. Comprehension prompts
    Ask for explanation of specific parts of the response

## Prompting Skills

### Ambiguity Reduction

Clearly specify the context and desired format of the response.
You can guide the model toward more valuable and precise outputs.

### Constraint-Based Prompting

Define explicit conditions or requirements in the prompt.
Encourages the model to produce a more focused and applicable response.

### Comparative Prompt Engineering

Ask the model to compare or contrast multiple entities or concepts.
Evaluates a models understanding

## Prompt templates

Predefined structures/formats that can be filled with specific content.

3 key steps

### Analyse the task

Understand the core requirements of the task at hand.
Determine specific info and concepts that need to be conveyed to elicit the desired outcome

### Design the structure

Create a template with placeholders to accommodate different inputs and situations.

Customer support example: "Hi, I'm experiencing [Issue/Problem] with [Product/Service]. Could you please help me resolve this? I've already tried [Actions Taken]."

I'm going to [location] with [group] for [time] please suggest an itinerary for the following interests [interests]

### Test and refine

Try out the template, based on the response refine it.

## Innovative prompt options

### Delimiters

Symbolic cues that aid in structuring ChatGPT responses.
Any symbols can be used as delimiters, you need to instruct the AI what to do with the text between them.
Example:
* Replace the text between them with emojis
* Pseudo code boolean conditions
* text between cues becomes section names

Delimiters can also prevent text from being taken up as an instruction.
For example if you have a chatbot on your site that is talking to an LLM in the background with user provided input.
The user could try and sneak instructions in to screw up what it returned by the LLM when you give it the user input and your instruction

### Requesting diverse output formats

Without specifying it ChatGPT will pick an output format.
You can specify what output format you want it in, numbered list, dialogue, Q&A, json, html, xml!
You can also specify the keys in in json to help your backend webserver easily deserialize the response.

### Translation services

ChatGPT exhibits remarkable translation capabilities.
Provide text in one language and request translation to another.

### Writing and summarization

Excels at condensing complex concepts and summarizing extensive text.

### Sentient analysis and adjustment

ChatGPT can identify the Sentient in your text and even modify the tone, output based on that.

### Inference capabilities

Demonstrates the ability to infer various elements such as central theme or the author's intent.

### Brainstorming assistance

Can facilitate creative process by aiding in brainstorming

### Project Development

Assist in strategizing and structuring your project.
help with the first few steps of starting a business

## Priming

Set the scene for the AI, then employ prompts to get the responses.

Primer: "You are a [] who has []. You possess []"

Prompt: "With this background, ..."

### 4 keys to effective priming

Granular Persona Details
  age, experience, industry, motivations, world view
Required Knowledge Breadth
  technical skills, industry, type of analysis, creative medium
Conversation Expectations
  professionalism, length, structure, jargon
Scenario Modelling
  constraints, priorty elements to feature, ...
Applicable Business Use Cases

### Advanced Prompt Engineering

Role-Driven Persona Contextualization
  give the AI a precise point of view
Multi-Layered Prompt Directives
  break down a complex goal into chunks, which the AI can figure out individually
  <Prompt> Then base on X do Y. Finally, ...
Hypothetical Scenario Simulation
  Use phrases like "Imagine if" or "In the year 2030"
Constraint-Bound Complex Inquiry
  Constrain the output of the AI
  simple language, limit text, examples or bullet points
Iterative Refinement Loops
  repeat prompts with evolving criteria
Meta-Prompting
  Get the AI to generate prompts
Nuanced Negative Prompting
  guide response by limiting input data

## frameworks

### A.P.E

Action Purpose Expectation.
  Define the job or activity to be done
          Discuss the intention or goal
                  State the desired outcome

### R.A.C.E

Role Action Context Expectations
  Specify the role of ChatGPT
      Detail what action is needed
              Provide relevant details of the situation
                        Describe the expected outcome

### C.O.A.S.T.

Context Objective Actions Scenario Task
  Set the stage for the conversation
          describe the goal
                    Explain the actions needed
                            Describe the scenario
                                    describe the task

### T.A.G.

Task Action Goal
  Define the specific task
        Describe what needs to be done
              Explain the end goal

### R.I.S.E

Role Input Steps Expectation
  Specify the role of ChatGPT
      Describe the info or resource
            Ask for detailed Steps
                    Describe the desired result
