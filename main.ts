/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-12-11
 * @fileoverview This program will compleate the mathmatical operation requested by the user
 */

console.log("Select an operation(type the letter beside the mathmatical operation):");
console.log("a. add");
console.log("b. subtract");
console.log("c. multiply");
console.log("d. divide");
console.log("e. absolute value");
console.log("f. round");
console.log("g. raise to an exponent");
console.log("h. square root");

let choice: string | null = prompt("Enter choice: ");
if (choice) {
  choice = choice.toLowerCase().trim();
}

let num1: number;
let num2: number;

if (choice === "a") { // add
  num1 = Number(prompt("Enter first number: "));
  num2 = Number(prompt("Enter second number: "));
  console.log(`${num1} + ${num2} = ${num1 + num2}`);
} else if (choice === "b") { // subtract
  num1 = Number(prompt("Enter first number: "));
  num2 = Number(prompt("Enter second number: "));
  console.log(`${num1} - ${num2} = ${num1 - num2}`);
} else if (choice === "c") { // multiply
  num1 = Number(prompt("Enter first number: "));
  num2 = Number(prompt("Enter second number: "));
  console.log(`${num1} * ${num2} = ${num1 * num2}`);
} else if (choice === "d") { // divide
  num1 = Number(prompt("Enter first number: "));
  num2 = Number(prompt("Enter second number: "));
  console.log(`${num1} / ${num2} = ${num1 / num2}`);
} else if (choice === "e") { // absolute value
  num1 = Number(prompt("Enter number: "));
  console.log(`|${num1}| = ${Math.abs(num1)}`);
} else if (choice === "f") { // round
  num1 = Number(prompt("Enter number: "));
  console.log(`round(${num1}) = ${Math.round(num1)}`);
} else if (choice === "g") { // exponent
  num1 = Number(prompt("Enter base: "));
  num2 = Number(prompt("Enter exponent: "));
  console.log(`${num1} ^ ${num2} = ${Math.pow(num1, num2)}`);
} else if (choice === "h") { // square root
  num1 = Number(prompt("Enter number: "));
  console.log(`sqrt(${num1}) = ${Math.sqrt(num1)}`);
} else {
  console.log("Invalid choice.");
}
console.log("\nDone.")