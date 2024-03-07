# Wheel of Fortune

This was quick tool built to solve the most important problem for every community radio station! 

*Problem: Who should engineer the big show?* 

### How to use 

Clone the project locally 

```
git clone git@github.com:shreypuranik/wheel-of-fortune.git
```

Set up the participants text file: 

```
$ make create-participants 
```

This will create `participants.txt`. Populate it with the relevant names. 

Run the builder

```
$ make build 
```
This will package up the project 

Spin the wheel! 

```
$ make spin 
```

This will tell you who has won!

```
➜  $ make spin
./wheel-of-fortune
Spinning the wheel. Who will engineer todays show?

The winner is Steve
```
Have the best show! :) 



