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
    INACTIVE=1,
    ACTIVE ,
    COMPLETED,
    ARCHIVED,
}

export const projectStatusMap = {
    [ProjectStatus.INACTIVE]: "Inactive",
    [ProjectStatus.ACTIVE]: "Active",
    [ProjectStatus.COMPLETED]: "Completed",
    [ProjectStatus.ARCHIVED]: "Archived",
}

