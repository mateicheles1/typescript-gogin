export class Album {
  id: string;
  title: string;
  artist: string;
  price: string;

  constructor(data: any) {
    this.id = data.id as string;
    this.title = data.title;
    this.artist = data.artist;
    this.price = data.price as string;
  }

  printData(
    parentComponent: HTMLElement | null,
    childComponent: HTMLDivElement
  ) {
    parentComponent?.appendChild(childComponent);
  }

  createDivChildComponent(): HTMLDivElement {
    const el = document.createElement("div");
    el.classList.add("album");
    return el;
  }
}
