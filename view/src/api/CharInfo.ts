export type CharInfo = {
    char: string,
    pos: {
        x1: number, y1: number, x2: number, y2: number,
    },
    display: boolean,
    selected: boolean,
    saved: boolean,
}