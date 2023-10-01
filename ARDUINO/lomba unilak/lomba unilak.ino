#include <LiquidCrystal_I2C.h>
LiquidCrystal_I2C lcd(0x27, 16, 2);

#include <SoftwareSerial.h>
#include <DFPlayer_Mini_Mp3.h>

SoftwareSerial serial(10, 11);

//  Variables
int PulseSensorPurplePin = 0;        // Pulse Sensor PURPLE WIRE connected to ANALOG PIN 0
int VIBRATOR1 = 12;   //  The on-board Arduion LED
int VIBRATOR2 = 7;

int Signal;    
int BPM;
// holds the incoming raw data. Signal value can range from 0-1024
int Threshold = 60;            // Determine which Signal to "count as a beat", and which to ingore.


// The SetUp Function:
void setup() {
   lcd.begin();
  lcd.backlight();
  pinMode(VIBRATOR1,OUTPUT); 
   pinMode(VIBRATOR2,OUTPUT); // pin that will blink to your heartbeat!
   Serial.begin(9600);         // Set's up Serial Communication at certain speed.
  serial.begin (9600);  
  mp3_set_serial (serial);
  delay(5); 
  mp3_set_volume (50);
  delay(1000);
}

// The Main Loop Function
void loop() {

  Signal = analogRead(PulseSensorPurplePin);  // Read the PulseSensor's value.
  BPM = Signal*0.15;                                            // Assign this value to the "Signal" variable.
Serial.println("BPM :");
   Serial.println(BPM);                    // Send the Signal value to Serial Plotter.
 lcd.setCursor(0,0);
lcd.println("BPM : ");
lcd.println(BPM);
delay(1000);

   if(BPM > 60){                          // If the signal is above "550", then "turn-on" Arduino's on-Board LED.
      lcd.setCursor(0,1);
lcd.println("NORMAL");
     digitalWrite(VIBRATOR1,LOW);
     digitalWrite(VIBRATOR2,LOW);
     delay(1000);
   } else {
              lcd.setCursor(0,1);
lcd.println("AWAS KANTUK");
     digitalWrite(VIBRATOR1,HIGH);
     digitalWrite(VIBRATOR2,HIGH);               //  Else, the sigal must be below "550", so "turn-off" this LED.
     mp3_play(1);
  delay(33120);
   }


delay(10);


}