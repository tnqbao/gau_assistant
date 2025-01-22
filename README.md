# AI Assistant using Gemini-1.5-Flash

This is a simple chatbot that uses the Gemini-1.5-Flash model to generate responses to user queries. The chatbot is
built using the Rasa framework and the Gemini-1.5-Flash model. It can answer questions, provide information, and engage
in conversations with users. The chatbot is designed to be user-friendly and customizable to meet the needs of different
users.

## Platforms Intergrated

<ul>
    <li>Messenger</li>
    <li>Discord</li>
    <li>Update soon ...</li>
</ul>

## Prerequisites

<ul>
  <li>Go 1.19 or newer</li>
  <li><a href="https://www.postman.com/" target="_blank">Postman</a> (optional, for API testing)</li>
</ul>

## Installation

* Clone the repository:
   ``` bash
    git clone https://github.com/tnqbao/gau_assistant.git
    cd gau_assistant
   ```
* Setup your module:
  ``` bash
   go mod edit -module=your-link-github-repo 
  ```
* Install dependencies:
  ``` bash
    go mod tidy 
  ``` 

* Set up environment variables:
    * Create a `.env` file in the project root and configure the following variables:
  ```dotenv
    GEMINI_API_KEY=your-gemini-api-key
    GEMINI_API_URL=api-model-url
    DISCORD_TOKEN=your-discord-token
    MESSENGER_TOKEN=your-messenger-token
    ```


* Start the server:
    ``` bash 
    go run main.go
    ```