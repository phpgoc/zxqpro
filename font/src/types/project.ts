export enum RoleType {
    OWNER = 1,
    PRODUCTER,
    DEVELOPER,
    TESTER,
    ADMIN 
}

export const roleTypesMap = {
    [RoleType.OWNER]: "Owner",
    [RoleType.PRODUCTER]: "Producter",
    [RoleType.DEVELOPER]: "Developer",
    [RoleType.TESTER]: "Tester",
    [RoleType.ADMIN]: "Admin"
}

export enum ProjectStatus {
    ACTIVE = 1,
    INACTIVE,
    COMPLETED,
    ARCHIVED,
}

export const projectStatusMap = {
    [ProjectStatus.ACTIVE]: "Active",
    [ProjectStatus.INACTIVE]: "Inactive",
    [ProjectStatus.COMPLETED]: "Completed",
    [ProjectStatus.ARCHIVED]: "Archived",
}