# Ascii-art-fs

This is a go project that receives a string as an argument and outputs the string in a graphical representation using ASCII art characters. The program ensures the correct representation of each character, adhering to ASCII standards.

# Features

1. **[ContainNonPrintChar](https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs/src/branch/master/ascii/nonprint.go):** Ascii has printable and non-printable characters, this function ensures that the non-printable characters are excluded in this project.
2. **[ReplaceNonPrintChar](https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs/src/branch/master/ascii/replace.go):** This function ensures that the escape sequences are replaced with their printable representations.
3. **[SpecialCases](https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs/src/branch/master/ascii/specialcase.go):** This function allows the execution of new line and empty string.
4. **[PrintChar](https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs/src/branch/master/ascii/printchar.go):** This function reads the banner files and gives the output of the file the way it appears. Each character takes a maximum of 8 lines in the banner files.
5. **[ProcessWords](https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs/src/branch/master/ascii/processwords.go):** This function processes the input and ensures that the output is printed according to the banner file stated.
6. Four types of ASCII art are available:

 **Standard**:
```bash
 _    _  $
| |  | | $
| |__| | $
|  __  | $
| |  | | $
|_|  |_| $
         $
         $
```
 **Shadow**:
```bash
         $
_|    _| $
_|    _| $
_|_|_|_| $
_|    _| $
_|    _| $
         $
         $
```


 **Thinkertoy**:
```bash
     $
o  o $
|  | $
O--O $
|  | $
o  o $
     $
     $
```

 **Zero**:
```

  0   $
 00   $
0 0   $
  0   $
  0   $
  0   $
00000 $
      $
```

**Note** 

* Program handles input with numbers, letters, spaces, special characters and \n.
* Based on the flag, program gives the user an opportunity to select which ASCII art file to use. 
* Program uses flags;  "shadow", "thinkertoy", "standard", "zero" for "shadow.txt", "thinkertoy.txt", "standard.txt" and zero.txt files respectively.  
* Additionally, the program will still be able to run with a single [STRING] argument and it will use standard.txt file to retrieve the ascii art characters. 



# Installation

Make sure you have Go installed on your machine.

1. Clone the repository

```bash
git clone https://learn.zone01kisumu.ke/git/somotto/ascii-art-fs.git
```

2. Navigate to the project repository

``` bash
cd ascii-art-fs
```

3. Run the project

```bash
go run . "hello" standard | cat -e
```

## Usage
The usage respect the format go run . [STRING] and go run . [STRING] [BANNER], any other formats returns the following usage message:

```bash
go run . [STRING] [BANNER]

go run . something standard 
```

 **Standard.txt**

* Input

``` bash
go run . "Hello\nThere" standard | cat -e
```

* Output

```bash
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
 _______   _                            $
|__   __| | |                           $
   | |    | |__     ___  _ __    ___    $
   | |    |  _ \  / _ \ | '__|  / _  \  $
   | |    | | | | |  __/ | |    |  __/  $
   |_|    |_| |_|  \___| |_|     \___|  $
                                        $
                                        $

```
## License:

This project is licensed under the MIT License. See the [LICENSE](https://en.wikipedia.org/wiki/MIT_License) file for more details.

## Testing and How to Run Tests:

Tests for the project are available in the printchar_test.go file. To run the tests, follow these steps:

* Navigate to the project directory.
```
cd ascii/
```
* Run the following command:
```
go test -v
```
This will run all the tests in the directory(ascii) and display the results.

## Contribution

Contributions to the project are welcome. If you'd like to contribute, follow these steps:

*   Fork the repository.
*   Create a new branch for your feature or bug fix.
*   Make your changes and commit them with clear, descriptive messages.
*   Push your changes to your fork.
*   Submit a pull request to the main repository.


## Collaborators
 This project was made possible from the contributions made by the following team members: 
*  [somotto](https://learn.zone01kisumu.ke/git/somotto)
*  [doonyango](https://learn.zone01kisumu.ke/git/doonyango)
