import { Album } from "../models/models";
import { Container } from "../constants/constants";

async function getData(): Promise<Album[]> {
  const response = await fetch("http://localhost:8080/albums");
  const data: Album[] = await response.json();
  return data.map((album: Album) => new Album(album));
}

export async function PrintData() {
  const data = await getData();
  data.forEach((album: Album) => {
    const el = document.createElement("div");
    el.classList.add("album");
    el.setAttribute;
    for (let keys in album) {
      el.textContent += `${keys}: ${album[keys]} `;
      Container?.appendChild(el);
    }
  });
}
