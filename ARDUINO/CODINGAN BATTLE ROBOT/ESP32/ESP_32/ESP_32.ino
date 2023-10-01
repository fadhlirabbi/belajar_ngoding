//2 SERVO GRIPPER.................(full arduino pro mini)
  #include <Servo.h>       // library servo
    Servo servoX;          // untuk servo yang pencapit 
    Servo servoY;          // untuk servo yang naik atau turun
  int pos = 0; 
  
//BLUETOOTH CONNECTION 
  #include <BluetoothSerial.h>

//SERIAL KOMUNIKASI 
  #include <SoftwareSerial.h> 
  SoftwareSerial serial_com(3,2);  // untuk pin 3 sebagai Rx dan pin 2 sebagai Tx
  String data;

//INFRARED SENSOR (IN ESP32).......(full esp32)
  #define S1  23
  #define S2  22
  #define S3  1
  #define S4  3
  #define S5  21
  #define S6  19
  #define S7  18
  #define S8  16
  #define S9  4 
  #define S10 2
  #define S11 15
  #define S12 32

//TB6612FNG..........................(Full Esp32)
  #include <SparkFun_TB6612.h>
  #define AIN1 3
  #define BIN1 7
  #define AIN2 4
  #define BIN2 8
  #define PWMA 5
  #define PWMB 6
  #define STBY 9

  // these constants are used to allow you to make your motor configuration 
  // line up with function names like forward.  Value can be 1 or -1
  const int offsetA = 1;
  const int offsetB = 1;

  // Initializing motors.  The library will allow you to initialize as many
  // motors as you have memory for.  If you are using functions like forward
  // that take 2 motors as arguements you can either write new functions or
  // call the function more than once.
  Motor motor1 = Motor(AIN1, AIN2, PWMA, offsetA, STBY);
  Motor motor2 = Motor(BIN1, BIN2, PWMB, offsetB, STBY);
 
//SENSOR WARNA TCS230...............(campuran esp & arduino pro mini)
  #include <MD_TCS230.h>
  #include <FreqCount.h>
  #define S0 4
  #define S1 5
  #define S2 6
  #define S3 7
  #define sensorOut 8
  
  // Variables for Color Pulse Width Measurements
  int redPW = 0;
  int greenPW = 0;
  int bluePW = 0;


void setup() {
  
//2 SERVO GRIPPER
  servoX.attach(6); 
  servoY.attach(7);

// setup tcs230
  // Set S0 - S3 as outputs
  pinMode(S0, OUTPUT);
  pinMode(S1, OUTPUT);
  pinMode(S2, OUTPUT);
  pinMode(S3, OUTPUT);
  
  // Set Pulse Width scaling to 20%
  digitalWrite(S0,HIGH);
  digitalWrite(S1,LOW);

  // Set Sensor output as input
  pinMode(sensorOut, INPUT);

//BLUETOOTH SETUP 
  BTSerial.write("AT\r\n");                       // expect "OK"
  BTSerial.setPinCode("1234");                    //PASSWORD FOR CONNECT THE BLUETOOTH
  BTSerial.setDeviceName("ESP32_BATTLE_ROBOT");   //NAMA PERANGKAT BLUETOOTH 

//SERIAL KOMUNIKASI 
  Serial.begin(9600);
}

void loop() {
  
//TCS230
 
// Read Red Pulse Width
  redPW = getRedPW();
  // Delay to stabilize sensor
  delay(200);

  // Read Green Pulse Width
  greenPW = getGreenPW();
  // Delay to stabilize sensor
  delay(200);

  // Read Blue Pulse Width
  bluePW = getBluePW();
  // Delay to stabilize sensor
  delay(200);

  // Print output to Serial Monitor
  Serial.print("Red PW = ");
  Serial.print(redPW);
  Serial.print(" - Green PW = ");
  Serial.print(greenPW);
  Serial.print(" - Blue PW = ");
  Serial.println(bluePW);  

  // Function to read Red Pulse Widths
  int getRedPW() {
      
    // Set sensor to read Red only
    digitalWrite(S2,LOW);
    digitalWrite(S3,LOW);
    
    // Define integer to represent Pulse Width
    int PW;
    
    // Read the output Pulse Width
    PW = pulseIn(sensorOut, LOW);
    // Return the value
    return PW;
    }
  
  // Function to read Green Pulse Widths
  int getGreenPW() {
    // Set sensor to read Green only
    digitalWrite(S2,HIGH);
    digitalWrite(S3,HIGH);
    
    // Define integer to represent Pulse Width
    int PW;
    
    // Read the output Pulse Width
    PW = pulseIn(sensorOut, LOW);
    // Return the value
    return PW;
    }
  
  // Function to read Blue Pulse Widths
  int getBluePW() {
    
    // Set sensor to read Blue only
    digitalWrite(S2,LOW);
    digitalWrite(S3,HIGH);
    
    // Define integer to represent Pulse Width
    int PW;
    
    // Read the output Pulse Width
    PW = pulseIn(sensorOut, LOW);
    // Return the value
    return PW;
    }

//tb6612fng 

//Use of the drive function which takes as arguements the speed
   //and optional duration.  A negative speed will cause it to go
   //backwards.  Speed can be from -255 to 255.  Also use of the 
   //brake function which takes no arguements.
   motor1.drive(255,1000);
   motor1.drive(-255,1000);
   motor1.brake();
   delay(1000);
   
   //Use of the drive function which takes as arguements the speed
   //and optional duration.  A negative speed will cause it to go
   //backwards.  Speed can be from -255 to 255.  Also use of the 
   //brake function which takes no arguements.
   motor2.drive(255,1000);
   motor2.drive(-255,1000);
   motor2.brake();
   delay(1000);
   
   //Use of the forward function, which takes as arguements two motors
   //and optionally a speed.  If a negative number is used for speed
   //it will go backwards
   forward(motor1, motor2, 150);
   delay(1000);
   
   //Use of the back function, which takes as arguments two motors 
   //and optionally a speed.  Either a positive number or a negative
   //number for speed will cause it to go backwards
   back(motor1, motor2, -150);
   delay(1000);
   
   //Use of the brake function which takes as arguments two motors.
   //Note that functions do not stop motors on their own.
   brake(motor1, motor2);
   delay(1000);
   
   //Use of the left and right functions which take as arguements two
   //motors and a speed.  This function turns both motors to move in 
   //the appropriate direction.  For turning a single motor use drive.
   left(motor1, motor2, 100);
   delay(1000);
   right(motor1, motor2, 100);
   delay(1000);
   
   //Use of brake again.
   brake(motor1, motor2);
   delay(1000);
  }
