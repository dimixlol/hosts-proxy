export class HostValidator implements validator {
    origin: string
    regex = new RegExp("^(?!:\\/\\/)(?=.{1,255}$)((.{1,63}\\.){1,127}(?![0-9]*$)[a-z0-9-]+\\.?)$")
    constructor(origin: string) {
        this.origin = origin
    }
    public validate(): boolean {
      return this.regex.test(this.origin)
    }
}
