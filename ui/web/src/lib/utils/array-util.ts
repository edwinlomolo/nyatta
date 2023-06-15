export function createArrayWithNumber(size: number): Array<number> {
  return Array.from({ length: size }, (_, i) => i+1)
}
