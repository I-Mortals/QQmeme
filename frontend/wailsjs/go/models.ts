export namespace memeFile {
	
	export class MemeInfo {
	    Name: string;
	    Code: string;
	    ParentPath: string;
	    Icon: string;
	    Memes: string[];
	
	    static createFrom(source: any = {}) {
	        return new MemeInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Code = source["Code"];
	        this.ParentPath = source["ParentPath"];
	        this.Icon = source["Icon"];
	        this.Memes = source["Memes"];
	    }
	}

}

