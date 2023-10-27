  

# LC-random
This is a CLI App built using Golang. It gives you random Leetcode questions to solve each time your run it. You can filter by topic as well and add your own questions from any platform.

## Examples 
### Mac using iTerm Terminal 
<img src="https://github.com/sameerkhan24/LC-random/assets/68337661/0c642002-02d0-4b1e-a63e-7746b941710c" alt="Image" width="600" />

### Windows using Command Prompt 
<img src="https://github.com/sameerkhan24/LC-random/assets/68337661/1053b155-727a-48e9-85d3-d583a240a36d" alt="Image" width="600" />

### Excel Template
  ![image](https://github.com/sameerkhan24/LC-random/assets/68337661/46c6ca62-e793-44f5-9709-dc2fbb2db710)


## Usage
Executable files and Excel template are available in `USE` folder. There are separate folders for Linus, windows and Mac. There is also a zip file also with Windows and Mac executables and LC template. 

### Windows 

 - Download `LC-random.exe` and `LC.xlsx` from
   [LC-random/USE/Windows](https://github.com/sameerkhan24/LC-random/tree/main/USE/Windows)
- Navigate to the folder where its located and open **Command Prompt**
- Execute the command `LC-random` and it will give you a random Leetcode problem with the problem link that can be clicked by holding CTRL and clicking on it 
- You can filter by topic with the `-topic` flag. 
E.g., `LC-random -topic "Sliding Window"`
- You can change the name of the LC-random.exe file and run commands with new name if you want
- If you will reply **Yes** to the question `Did you solve it? (Yes/No)` it will update the **Times Solved** count by 1 in `LC.xlsx`. Count will not be updated if you reply **No**

### Mac

- It is preferred to use iTerm, VS code or any other terminal to execute this as the default terminal on Mac OS doesn't allow links to be clickable
 - Download `LC-random` and `LC.xlsx` from
   [LC-random/USE/Mac](https://github.com/sameerkhan24/LC-random/tree/main/USE/Mac)
- Navigate to the folder where its located and open the **iTerm/vscode**
- Run the command `chmod +x LC-random` to make it executable
- Run the command `./LC-random` and it will give you a random Leetcode problem with the problem link that can be clicked by holding Command and clicking on it 
- You can filter by topic with the `-topic` flag. 
E.g., `./LC-random -topic "Sliding Window"`
- You can change the name of the LC-random file and run commands with new name if you want
-  If you will reply **Yes** to the question `Did you solve it? (Yes/No)` it will update the **Times Solved** count by 1 in `LC.xlsx`. Count will not be updated if you reply **No**


## Points to remember 

 - You can change the name of the executable files. The executing command will be changed accordingly
 - You can't change the name of the Excel file from LC.xlsx. 
 - LC.xlsx should be in the same location as the executable file
 - You can edit LC.xlsx to add more questions
 - The columns for PROBLEM, TOPIC, LINK, and Times Solved should not be removed or tampered with. 
 - Make sure to add the link is when you add new problems

