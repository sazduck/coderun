// @ts-check
// Для чтения входных данных в Node.js необходимо использовать
// модуль readline, который работает с потоком ввода-вывода
// (stdin/stdout) и позволяет читать строки.
import { createInterface } from 'readline';
const rl = createInterface({
  input: process.stdin,
  output: process.stdout,
});

// Считывание данных из нескольких строк в переменную
/** @type {string[]}*/
const input = [];
rl.on('line', line => {
  input.push(line);
});

// Код решения можно писать внутри функции закрытия потока ввода
rl.on('close', () => {
  const wordSet = new Set();
  input
    .map(line => line.split(' ').filter(line => line))
    .flat()
    .forEach(word => wordSet.add(word));
  console.log(wordSet.size);
});
