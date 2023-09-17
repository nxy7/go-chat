//@ts-ignore
import { expect, test } from "bun:test";
import { DecodeJwt } from "./userStore";

test("jwt decoding", () => {
    let jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJleHAiOjE1MTYyMzkwMjJ9.mVGXFv3OuwtuZPsdaf_oGUYm2uOH-T-JRTDQE1c10q0"
    let decoded = DecodeJwt(jwt)
    console.log(decoded)
    // expect(2 + 2).toBe(4);
});