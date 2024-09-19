export function decodedValue(colors: string[]): number {
  let index = -1
  let counter = 0
  let value = 0

  for (let color of colors) {
    index = COLORS.findIndex(c => c === color.toLowerCase())
    counter++;
    if (index === -1){
      throw new Error('invalid color: ' + color)
    }
    if (counter ===2) {
      value += index
    }else if (counter === 1) {
      value = index*10
    } 
  }

  return value
}


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
