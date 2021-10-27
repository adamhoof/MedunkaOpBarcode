# MedunkaOpBarcode
#### The main goal is to optimize and speed up goods unpacking by not having to put price tags on products, but still being able to show product info to customers.

- The brain -> RPI Zero W.
- Program function -> Written in Go. After scanning barcode, either query of local database with product data is done, or request is sent to API endpoint, displayed to user.
- Barcode reading -> Waveshare barcode reader. Using UART interface, since transmission via USB resulted in malformed output, even tho it worked with other devices.
- Product data display -> Waveshare display with HDMI interface.
- RPI and logic level converter for RPI<-->BarcodeScanner communication is placed in 3D printed box.
- Barcode scanner is also placed in a 3D printed box, with taper angle facing down to avoid eye damage to users.
