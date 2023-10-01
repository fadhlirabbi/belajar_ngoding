#include "Arduino.h"
#include "SoftwareSerial.h"      // serial
#include "DFRobotDFPlayerMini.h" // dfplayer

//#include <SPI.h>
#include <Wire.h> 
#include <LiquidCrystal_I2C.h>    // lcd
LiquidCrystal_I2C lcd(0x27, 16, 2);

#include <Adafruit_GFX.h>
//#include <DFPlayer_Mini_Mp3.h> 

int stats=0;
SoftwareSerial mySoftwareSerial(10, 11); // RX, TX
DFRobotDFPlayerMini myDFPlayer;
void printDetail(uint8_t type, int value);
String readString;
int g,j,x;
int myBPM;

unsigned long myTime;
unsigned long time_counting;
unsigned long time_threshold = 7000;
const int PulseWire = 0;
const int vibrator_1 = 12;
const int vibrator_2 = 7;

int i=0;

void setup() {
  mySoftwareSerial.begin(9600);
  Serial.begin(115200);
 
  Serial.println(F("DFPlayer Mini online."));
  myDFPlayer.volume(20);  //Set volume value. From 0 to 30
  
  pinMode(vibrator_1,OUTPUT);
  pinMode(vibrator_2,OUTPUT);
  Serial.println("Input BPM Number"); // so I can keep track of what is loaded
} 

void vibrator_on(){
  Serial.println("Vibrator ON");
  digitalWrite(vibrator_1,HIGH);
  digitalWrite(vibrator_2,HIGH);

}

void vibrator_off(){
  Serial.println("Vibrator OFF");
  digitalWrite(vibrator_1,LOW);
  digitalWrite(vibrator_2,LOW);  
}


void loop() {
  while (Serial.available()) {
    g=1;
    char c = Serial.read();  //gets one byte from serial buffer
    readString += c; //makes the string readString
    delay(2);  //slow looping to allow buffer to fill with next character
  }
  
  if (readString.length() >0 || g==1) {
    //Serial.println(readString);  //so you can see the captured string 
    int n = readString.toInt();  //convert readString into a number
    j++;
    x=1;
    if(j==1){
      myBPM = n;
    }
    if(n!=myBPM && n>0 && (myBPM>0 && myBPM<150)){
      myBPM = n;
    }
    readString=""; //empty for next input
  }
  
  if(x==1){
    myTime = millis();
    Serial.println(String() + "Nilai BPM Sekarang = " + myBPM);
    
    // Oled Display Program 
    lcd.clear(); // mengosongkan tampilan
    lcd.setCursor(0,0); // sett posisi tampilan
    lcd.println("BPM = " + String(myBPM)); // mengisi karakter
    
    if (myBPM < 60 || myBPM > 90){
      vibrator_on();
      if(stats == 0){
        lcd.setCursor(0,1); // sett posisi tampilan
        lcd.println("BPM RENDAH"); // mengisi karakter
      }
      //Serial.println("BPM Tidak Normal");
      if(myTime - time_counting >= time_threshold){
        Serial.println("Microsleep Terdeteksi");
        stats=1;
        if(stats==1){
          lcd.setCursor(0,1); // sett posisi tampilan
          lcd.println("MICROSLEEP"); // mengisi karakter
        }

        if (i==0){
          i = 1;
          myDFPlayer.play(random(1,4));  //Play the first mp3
        }
      }
    }
    else{
      stats=0;
      vibrator_off();
      i=0;
      time_counting = myTime;
      Serial.println(String() + "BPM NORMAL dan Waktu = " + time_counting);
      myDFPlayer.stop();
      lcd.setCursor(0,1); // sett posisi tampilan
      lcd.println("BPM NORMAL"); // mengisi karakter
    }
   lcd.display();
  }
  delay(100);
}



