interface FileChunkDataCallback {
    (data: Uint8Array): void
}

interface ErrorCallback {
    (message: string): void
}

interface FileReaderOnLoadCallback {
    (event: ProgressEvent): void
}

interface ProcessingCompletedCallback {
    (): void
}

class Validate {
    public static notNull<T>(obj: T): obj is Exclude<T, null> {
        if (obj == null) {
            throw new ReferenceError(`Error: Object ${obj} was null`);
        }
        return true;
    }

    public static notUndefined<T>(obj: T): obj is Exclude<T, undefined> {
        if (obj === undefined) {
            throw new ReferenceError(`Error: Object ${obj} was undefined`);
        }
        return true;
    }

    public static notNullNotUndefined<T>(obj: T): obj is NonNullable<T> {
        return Validate.notNull(obj) && Validate.notUndefined(obj);
    }
}

 // TODO use SubtleCrypto for small files, wasm fuck for large files

class FileInChunksProcessor {
    public readonly CHUNK_SIZE_IN_BYTES: number = 1024*1000*20;
    private readonly fileReader: FileReader;
    private readonly dataCallback: FileChunkDataCallback;
    private readonly errorCallback: ErrorCallback;
    private readonly processingCompletedCallback: ProcessingCompletedCallback;
    private start: number = 0;
    private end: number = this.start + this.CHUNK_SIZE_IN_BYTES;
    private inputFile: File | null = null;

    constructor(dataCallback: FileChunkDataCallback,
                errorCallback: ErrorCallback,
                processingCompletedCallback: ProcessingCompletedCallback) {
        this.fileReader = new FileReader();
        this.fileReader.onload = this.getFileReadOnLoadHandler();
        this.dataCallback = dataCallback;
        this.errorCallback = errorCallback;
        this.processingCompletedCallback = processingCompletedCallback;
    }

    public processChunks(inputFile: File) {
        this.inputFile = inputFile;
        this.read(this.start, this.end);
    }

    public getFileFromElement(elementId: string): File | undefined {
        const filesElement = document.getElementById(elementId) as HTMLInputElement;
        if (Validate.notNull(filesElement) &&
            Validate.notNull(filesElement.files)) {
            if (filesElement.files.length < 0) {
                this.errorCallback("Too few files specified. Please select one file.")
            } else if (filesElement.files.length > 1) {
                this.errorCallback("Too many files specified. Please select one file.")
            } else {
                return filesElement.files[0];
            }
        }
    }

    private getFileReadOnLoadHandler(): FileReaderOnLoadCallback {
        return () => {
            if (Validate.notNull(this.inputFile)) {
                this.dataCallback(new Uint8Array((this.fileReader.result as ArrayBuffer)));

                this.start = this.end;
                if (this.end < this.inputFile.size) {
                    this.end = this.start + this.CHUNK_SIZE_IN_BYTES;
                    // if(this.end > this.inputFile.size) {
                    //     this.end = this.inputFile.size;
                    // }
                    console.log(`loading ${this.start}...${this.end}`);
                    this.read(this.start, this.end);
                } else {
                    this.processingCompletedCallback();
                }
            }
        }
    }

    private read(start: number, end: number) {
        if (Validate.notNull(this.inputFile)) {
            this.fileReader.readAsArrayBuffer(this.inputFile.slice(start, end));
        }
    }
}

function handleError(message: string) {
    const errorElement = document.getElementById("error");
    if (Validate.notNull(errorElement)) {
        errorElement.innerHTML = `${errorElement.innerHTML} <p>${message}</p>`;
    }
}

function processFileButtonHandler() {
    const startElement = document.getElementById("start");
    if(Validate.notNull(startElement)) {
        startElement.innerText = "start " + time();
    }
    // @ts-ignore
    const hasher = CryptoJS.algo.SHA256.create();
    const processor = new FileInChunksProcessor((data) => {
            console.log('hashing..');
            // @ts-ignore
            hasher.update(CryptoJS.lib.WordArray.create(data));
        },
        handleError, () => {
            // @ts-ignore
            const hashStr = hasher.finalize().toString(CryptoJS.enc.Hex);
            const resultElement = document.getElementById("result");
            if (Validate.notNull(resultElement)) {
                resultElement.innerHTML = `<p>${time()} Got: ${hashStr} </p>`;
            }
        }
    );
    const file = processor.getFileFromElement("file");
    if (Validate.notNullNotUndefined(file)) {
        console.log(`Started at ${new Date().toLocaleTimeString()}`)
        processor.processChunks(file);
    }
}

function time() {
    const date = new Date();
    return date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds() + "." + date.getMilliseconds();
}
