# Assignment 21
Design a struct which will have members that describes an employee's:

* last name,
* first name,
* employee ID number
* last 4 digits of their NIN, and
* Title

Design a second struct which will store two employees as members. the employees will be structured of the type of your first struct. inside the main(), prompt the user to enter each employee's credentials and store them. finally, print the employee's credentials to the terminal

Pseudo Code:

type employee struct {
	members like: names, employee id, etc
}

type employees struct {
	employee1 employee
}

func min() {
	prompting and storing credentials for all employees
	printing all the data to the terminal
}

# Example Output
ftsog@nigeria:~/LearningGo/Assignment-21$ ./assignment21
Enter the employee's first name: Adam
Enter the employee's last name: Kolawole
Enter the employee's ID number: 345
Enter the last four digits of the employee's NIN: 9832
Enter the employee's job title (do not include the word 'Engineer'): Senior

Enter the employee's first name: Amanda
Enter the employee's last name: Nunes
Enter the employee's ID number: 678
Enter the last four digits of the employee's NIN: 4521
Enter the employee's job title (do not include the word 'Engineer'): Junior

Employee information for Adam Kolawole:
ID: 345
NIN: 9832
Title: Senior Engineer

Employee information for Amanda Nunes:
ID: 678
NIN: 4521
Title: Junior Engineer