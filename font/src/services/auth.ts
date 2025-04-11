export function isAdmin(userId: number, ownerId: number): boolean {
  return userId === ownerId || userId === 1;
}
