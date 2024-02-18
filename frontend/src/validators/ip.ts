abstract class ipAddress {
    protected _slug: string = ""
    static regex: RegExp = new RegExp("^$i")
    constructor(public address: string) { this.address = address }
    get slug(): Object {
        return {"slug":this._slug}
    }

}

export class IpV4Address extends ipAddress {
    protected _slug: string = "ipV4"
    static regex: RegExp = new RegExp("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")
}

export class IpValidator implements validator {
    origin: string
    private address?: ipAddress = undefined

    constructor(origin: string) {
        this.origin = origin
    }
    public validate(): boolean {
        if (IpV4Address.regex.test(this.origin)) {
            this.address = new IpV4Address(this.origin)
            return true
        }
        return false
    }
    public isIpV4(): boolean {
        return this.address instanceof IpV4Address
    }
}
