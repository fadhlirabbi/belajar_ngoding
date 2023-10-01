const int ledPin = 7;
const int ldrPin = A0;

void setup() {
Serial.begin(9600);
pinMode(ledPin, OUTPUT);
pinMode(ldrPin, INPUT);

}
void loop() {
    int ldrStatus = analogRead(ldrPin);

      if (ldrStatus > 200) {
      digitalWrite(ledPin, HIGH);
      Serial.print("LED ON, NILAI :");
      Serial.println(ldrStatus); 
      } else {
      digitalWrite(ledPin, LOW);
      Serial.print("LED OFF, NILAI :");
      Serial.println(ldrStatus);
      
      }
}
