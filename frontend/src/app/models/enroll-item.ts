import { UserInfo } from './user-info';
import { ProjectItem } from './project-item';

export interface EnrollItem {
    Id: number;
    User: UserInfo;
    Project: ProjectItem;
    Message: string;
    DateOfCreation: string;
}
