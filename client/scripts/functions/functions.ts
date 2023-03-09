import { Album } from "../models/models";

async function getData(): Promise<Album[]> {
  const response = await fetch("http://localhost:8080/albums");
  const data: Album[] = await response.json();
  return data.map((album: Album) => new Album(album));
}

async function printData() {
  const data = await getData();
}
