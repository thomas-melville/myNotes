# LLM

Large Language Model
A dataset that an AI is trained on.

## Types

Base LLMs - predicts next word based on training data
Instruction Tuned LLMs - tries to follow instructions and answer questions.
  starts as a Base LLM and is then trained on instructions
  RLHF: Reinforcement learning with Human Feedback

## Failure modes

Hallucinations - factually incorrect information generated by the AI
Bias - LLM is based on it's dataset, so if Bias is in the dataset it will be in the responses
Misinformation - Again based on the dataset, can't distinguish truth from falsehood.

## GPT

Generative Pretrained Transformers

Generative - it can create/generate responses based on it's data
Pretrained - previously trained on a vast amount of data.
Transformers - refers to the specific model that allows this to happen.

### Transformers model

Game changer when introduced by Google in 2017
It's all to do with how the AI handles the data.
Before it was word by word in a sentence, which was limited.
The Transformers model considers all words in a sentence at once.

#### Self improvement

To test it's understanding, the model sometimes hides a word in a sentence and tried to predict what the word is.


## API Access to LLM's

Use jupyter notebooks