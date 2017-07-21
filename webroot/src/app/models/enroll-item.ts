import { UserInfo } from './user-info';
import { ProjectItem } from './project-item';

export interface EnrollItem {
    User: UserInfo;
    Project: ProjectItem;
    Message: string;
    DateOfCreation: string;
}
