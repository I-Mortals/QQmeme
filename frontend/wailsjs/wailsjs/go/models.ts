export namespace memeFile {
	
	export class MemeInfo {
	    name: string;
	    code: string;
	    parentPath: string;
	    icon: string;
	    memes: string[];
	
	    static createFrom(source: any = {}) {
	        return new MemeInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.code = source["code"];
	        this.parentPath = source["parentPath"];
	        this.icon = source["icon"];
	        this.memes = source["memes"];
	    }
	}

}

