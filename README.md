GPT-4o-mini CLI

This is a command-line interface (CLI) for interacting with the OpenAI GPT-4o-mini model. It supports streaming responses and function calling.

## Installation

1.  **Install Go:** Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).
2.  **download the source code**, or clone the repository: git clone https://github.com/orender/gpt_cli_extented.git

## API Key Configuration

1.  **Create a `.env` file:** In the root directory of the project, create a file named `.env`.
2.  **Add your OpenAI API key:** Open the `.env` file and add your OpenAI API key as follows:
    ```
    OPENAI_API_KEY=your_api_key_here
    ```
    **Install the needed libraries

## Running the CLI

1.  **Run the Go program:**
    ```bash
    go run main.go
    ```
2.  **Interact with the CLI:** You will be prompted to enter your message. Type your message and press Enter. The CLI will stream the response from GPT-4o-mini.
3.  **Exit:** Type `exit` and press Enter to exit the CLI.

## Examples of Usage

* **Basic Conversation:**
    ```
    You:
    tell me the lore of dark souls

    GPT-4o-mini

    The lore of Dark Souls is intricate and deeply interconnected, encompassing themes of fire, darkness, sacrifice, and the cyclical nature of life and death. Here’s a brief overview of the key elements of the Dark Souls lore:

    1. **The Age of Ancients**: In the beginning, the world was unformed, shrouded in fog and inhabited by Everlasting Dragons. There was no fire, and everything was in a state of stagnation.

    2. **The First Flame**: The emergence of the First Flame marked a pivotal moment. It brought light to the world and introduced the concepts of disparity: light and dark, heat and cold, life and death. The First Flame also gave birth to powerful beings known as Lords.

    3. **The Lords**:
    - **Gwyn, Lord of Sunlight**: The leader of the gods who wielded the power of lightning and sunlight.
    - **Nito, the First of the Dead**: The lord of death and the grave, representing the cycle of life and mortality.
    - **The Witch of Izalith**: A powerful sorceress who attempted to create her own flame but instead caused chaos.
    - **Seath the Scaleless**: A dragon who betrayed his own kind and became a lord through his knowledge of sorcery.

    4. **The Undead Curse**: As time passed, a curse began to afflict humanity, leading to the rise of the Undead. Those who turned Undead lost their purpose and sanity, often spiraling into despair as they sought to break the curse.

    5. **The Linking of Fire**: To prolong the Age of Fire, the Lords decided to sacrifice themselves to the First Flame, resulting in the Linking of Fire, which perpetuated the existence of the flame but also led to the eventual fading of it.

    6. **The Age of Dark**: The Age of Fire cannot last forever, and with its fading comes the Age of Dark. Dark Souls explores the tension between these two ages and the cyclical nature of the world.

    7. **The Player’s Role**: Players assume the role of an Undead Chosen One, tasked with defeating powerful beings, reclaiming souls, and ultimately deciding the fate of the world—whether to link the fire and continue the Age of Fire or to let it die and usher in the Age of Dark.

    8. **Themes and Symbolism**: Dark Souls delves deeply into themes of despair, hope, sacrifice, and the nature of humanity. The world of Lordran is filled with rich lore, hidden secrets, and interconnected narratives, encouraging players to piece together the story through exploration and item descriptions.

    Dark Souls encourages speculation and interpretation, which adds to its allure and depth within the gaming community. Each player's journey through the game unfolds a unique story, shaped by their choices and experiences
    ```
* **Function Calling (Multiplication):**
    ```
    You:
    whats nine times nine

    GPT-4o-mini


    Function Name: multiply

    Result: 81


    You:
    whats 9 * 9

    GPT-4o-mini


    Function Name: multiply

    Result: 81


    ```

## Architecture and Design

The CLI is designed to provide a simple and efficient way to interact with the OpenAI GPT-4o-mini model. Here's a brief overview of the architecture:

* **Input Handling:** The `main` function uses `bufio.Scanner` to read user input from the command line.
* **API Request:** The `streamResponse` function constructs an HTTP POST request to the OpenAI API's chat completions endpoint. It sets the `Stream` parameter to `true` to enable streaming responses. The API key is loaded from the `.env` file using the `godotenv` package.
* **Streaming Response:** The `streamResponse` function uses `bufio.Reader` to read the streaming response line by line. It parses the JSON chunks and prints the content to the console.
* **Concurrency:** A goroutine and `sync.WaitGroup` are used to handle the streaming response concurrently, preventing the main thread from blocking. A context with timeout is used to prevent infinite loops.
* **Function Calling:** The CLI supports function calling. It parses the function call arguments from the streaming response and executes the corresponding function (in this example, multiplication).
* **Error Handling:** The code includes error handling for HTTP requests, JSON parsing, and file operations. Errors are logged using `log.Printf`.
* **Configuration:** The API key is configured using a `.env` file, which is loaded using the `godotenv` package.
* **Output:** The `fatih/color` package is used to colorize the output, making it more readable.
* **Maintainability:** The code is structured with clear functions and structs, making it easy to understand and maintain.
* **Extensibility:** The code can be easily extended to support other OpenAI models and functions.
* **Context:** a context with timeout is used to prevent the program from running infinitely if the API fails to respond.

This design ensures that the CLI is efficient, reliable, and easy to use.

**Note:** The GPT-4o-mini model's function calling behavior can be unpredictable at times. While improvements to the system prompt have been made, perfect consistency is not guaranteed.