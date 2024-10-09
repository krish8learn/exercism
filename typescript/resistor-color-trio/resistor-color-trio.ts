export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
] as const;

const units = ["", "kilo", "mega", "giga"];

export function decodedResistorValue(bands: string[]): string {
  // throw new Error('Remove this statement and implement this function')
  let index = -1
  let counter = 0
  let value = 0

  // find the value
  for (let color of bands) {
    let index = -1
    index = COLORS.findIndex(c => c === color.toLowerCase())
    counter++;
    if (index === -1){
      throw new Error('invalid color: ' + color)
    }
    if (counter === 1) {
      value = index*10
    }else if (counter === 2) {
      value += index
    }else if (counter === 3) {
      value = value * (10 ** index)
    }
  }


  // format the number in string
  return formatNumber(value)
}

function formatNumber (num : number): string {
  let unitIndex = 0;

    while (num >= 1000 && unitIndex < units.length - 1) {
        num /= 1000;
        unitIndex++;
    }

    return `${num} ${units[unitIndex]}ohms`;
}


