import { HOST_KEY } from "../types/const.ts";

export function ownerOrAdmin(userId: number, ownerId: number): boolean {
  return userId === ownerId || isAdmin(userId);
}

export function isAdmin(userId: number): boolean {
  return userId === 1;
}

export function serverUrl(): string {
  const webSetHost = localStorage.getItem(HOST_KEY) ;
  if(webSetHost) {
    return webSetHost as string;
  }
  return import.meta.env.VITE_SERVER_URL

}

export function avatarUrl(avatarId  : number): string {
  return `${serverUrl()}static/avatar/${avatarId}.webp`;
}

export function parseIdToNumber (id: string | undefined): number  {
  if (id) {
    const parsedId = Number(id);
    if (!isNaN(parsedId)) {
      return parsedId;
    }
  }
  return 0;
}