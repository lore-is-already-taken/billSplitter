import "./../../src/css/style.css";
import QRCode from "qrcode";

const message = "https://www.google.com/";
var canvas = document.getElementById("canvas");

QRCode.toCanvas(canvas, message, function (error) {
	if (error) console.error(error);
	console.log("success!");
});
