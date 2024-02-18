export interface Response<T> {
    code: number,
    message: string,
    ok: boolean,
    data: T,
}
