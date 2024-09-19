export const colorCode = (color: string) => {
  const index = COLORS.findIndex(c => c === color.toLowerCase());
  if (index === -1) {
    throw new Error(`Invalid color`);
  }
  return index;    
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
