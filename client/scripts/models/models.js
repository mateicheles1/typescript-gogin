"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Album = void 0;
class Album {
    constructor(data) {
        this.id = data.id;
        this.title = data.title;
        this.artist = data.artist;
        this.price = data.price;
    }
}
exports.Album = Album;
