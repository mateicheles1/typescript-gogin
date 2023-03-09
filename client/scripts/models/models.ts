export class Album {
  id: number;
  title: string;
  artist: string;
  price: number;
  constructor(data: any) {
    this.id = data.id;
    this.title = data.title;
    this.artist = data.artist;
    this.price = data.price;
  }
}
