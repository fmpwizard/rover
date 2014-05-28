// Adafruit Motor shield library
// copyright Adafruit Industries LLC, 2009
// this code is public domain, enjoy!

//Deploy this file to your arduino plugged into the rover

#include <AFMotor.h>

// DC motor on M1, Left wheel
AF_DCMotor left(1);
// DC motor on M2, Right wheel
AF_DCMotor right(2);

//The number of degrees to turn or how long to move for (in distance)
int value;
//Command, right, left, forward, backwards
String cmd;

void setup() {
  Serial.begin(9600);           // set up Serial library at 9600 bps
  Serial.println("Rover powered up!");
  
   
  // turn on motor #1
  right.setSpeed(100);
  left.setSpeed(100);
  right.run(RELEASE);
  left.run(RELEASE);
}


void loop() {
  if(Serial.available() > 2){
    cmd = Serial.readStringUntil(' ');
    value = Serial.parseInt() * 1000;
    processCommand();
  }
}

void processCommand() {
  
  if (cmd == "f") {
    right.run(FORWARD);
    left.run(FORWARD);
    delay(value);
    right.run(RELEASE);
    left.run(RELEASE);
  } else if (cmd == "b"){
    right.run(BACKWARD);
    left.run(BACKWARD);
    delay(value);
    right.run(RELEASE);
    left.run(RELEASE);
  } else if (cmd == "r"){
    right.run(BACKWARD);
    left.run(FORWARD);
    delay(value);
    right.run(RELEASE);
    left.run(RELEASE);
  } else if (cmd == "l"){
    right.run(FORWARD);
    left.run(BACKWARD);
    delay(value);
    right.run(RELEASE);
    left.run(RELEASE);
  }
}  