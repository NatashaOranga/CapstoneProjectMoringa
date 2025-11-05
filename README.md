# Bank Management System

A simple command-line bank management system built with Go (Golang) as part of a capstone project for learning a new programming language.

## Table of Contents
* [About](#about)
* [Features](#features)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
* [Project Structure](#project-structure)
* [How It Works](#how-it-works)
* [Future Enhancements](#future-enhancements)
* [Learning Resources](#learning-resources)
* [Technical Specifications](#technical-specifications)
* [Acknowledgments & Contact](#acknowledgments--contact)
* [License](#license)

---

## About

This project was created as a capstone for a short course focused on learning a new programming language. I chose **Go (Golang)** due to its simplicity, performance, and growing popularity in backend development and cloud infrastructure.

The Bank Management System demonstrates core Go concepts, including:
* Structs and methods with pointer receivers
* Maps for data storage (acting as an in-memory database)
* User input handling (`bufio` and `fmt`)
* Control flow (loops, conditionals, switch statements)
* Error handling

**Development Time:** 7 days (from zero Go knowledge to a working project).

## Features

The system includes the following core banking operations:

*  **Create Account** - Register new bank accounts with unique account numbers.
*  **Deposit Money** - Add funds to existing accounts.
*  **Withdraw Money** - Remove funds with balance validation.
*  **Check Balance** - View current account information.
*  **Input Validation** - Prevents invalid operations (negative amounts, insufficient funds).
*  **Interactive CLI** - User-friendly menu-driven interface.

## Prerequisites

Before running this project, ensure you have:

| Tool | Requirement | Verification |
| :--- | :--- | :--- |
| **Go** | Version 1.16 or higher installed | `go version` |
| **Text Editor/IDE** | Recommended: Visual Studio Code with Go extension | - |
| **Knowledge** | Basic command-line operations (navigate, run commands) | - |

## Installation

### Step 1: Clone or Download the Project
```bash
# Create a directory for the project (if you haven't already)
mkdir bank-management
cd bank-management

# Clone the repository
git clone [https://github.com/yourusername/bank-management.git](https://github.com/yourusername/bank-management.git) . 
````

*(Note: Replace `yourusername` with your actual GitHub username.)*

### Step 2: Ensure Go Module is Initialized

This step creates the required `go.mod` file.

```bash
go mod init bank-management
```

## Usage

### Running the Application

1.  Open your terminal/command prompt.
2.  Navigate to the project directory (`bank-management`).
3.  Run the program:
    ```bash
    go run main.go
    ```

### Using the System

Once the program starts, you'll be presented with the main menu:

```text
=== Bank Management System ===
1. Create Account
2. Deposit
3. Withdraw
4. Check Balance
5. Exit
Choose option:
```

**Example Workflow:**

1.  **Create an Account:** Choose `1`. Input Account Number, Holder Name, and Initial Balance (e.g., 1000).
2.  **Deposit Money:** Choose `2`. Input Account Number and Amount (e.g., 500).
3.  **Withdraw Money:** Choose `3`. Input Account Number and Amount (e.g., 200).
4.  **Check Balance:** Choose `4`. Input Account Number to view the current balance (e.g., 1300.00).

## Project Structure

```
bank-management/
├── main.go         # Main application code
├── go.mod          # Go module file
├── .gitignore      # Specifies files to ignore (like binaries)
└── LICENSE         # Defines usage rights
```

## How It Works

### Core Components

1.  **`Account` Struct:**

    ```go
    type Account struct {
        AccountNumber string
        HolderName    string
        Balance       float64
    }
    ```

    Represents the data structure for a single bank account.

2.  **Account Methods:**

      * **`Deposit` Method:** Uses a **pointer receiver** to modify the actual `Balance` in memory. Includes validation for positive amounts.
      * **`Withdraw` Method:** Uses a pointer receiver. Includes validation for both positive amounts and **sufficient funds**.
      * **`DisplayInfo` Method:** Formats and prints account details.

3.  **Data Storage:**

    ```go
    accounts := make(map[string]*Account)
    ```

    A Go map (dictionary) is used to store account pointers in memory, with the unique `AccountNumber` serving as the lookup key.

### Limitations (V1.0.0)

This is a basic implementation with the following limitations:

  * ⚠️ **No Data Persistence:** All data is lost when the program exits.
  * ⚠️ **In-Memory Storage:** Data is stored in RAM only.
  * ⚠️ **No Security:** No PIN or password protection.

## Future Enhancements

| Version | Planned Features |
| :--- | :--- |
| **V2.0** | Data persistence (save/load accounts from JSON files), Transaction history, Money transfer, PIN authentication. |
| **V3.0** | SQL database integration (PostgreSQL/MySQL), RESTful API instead of CLI, Concurrent user handling with **Goroutines**. |

## Learning Resources

### Resources Used During Development

  * [Official Go Documentation](https://go.dev/doc/)
  * [Go Tour](https://go.dev/tour/) - Interactive tutorial.
  * [Go by Example](https://gobyexample.com/) - Practical code examples.

### Building the Executable

To create a standalone executable for easier distribution:

```bash
# Build for your current OS
go build -o bank-system

# Run the executable
./bank-system
```

## Technical Specifications

| Specification | Details |
| :--- | :--- |
| **Language** | Go 1.21+ |
| **Dependencies** | Standard library only (`fmt`, `bufio`, `os`, `strconv`, `strings`) |
| **Project Type** | Capstone Learning Project (CLI) |

## Acknowledgments & Contact

  * **Go Team at Google:** For creating an amazing language.
  * **Course Instructors at Moringa:** For the learning framework.
  * **Go Community:** For excellent documentation and support.

| Contact Detail | Value |
| :--- | :--- |
| **Developer** | Natasha Oranga |
| **Email** | oranganatasha49@gmail.com |
| **Location** | Nairobi, Kenya |

## License

This project is open source and available under the **MIT License** for educational purposes.

-----

⭐ If you found this project helpful in learning Go, please star the repository\!

```
```# CapstoneProjectMoringa
