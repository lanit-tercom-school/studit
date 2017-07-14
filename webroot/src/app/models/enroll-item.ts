import { UserInfo } from './user-info';
import { ProjectItem } from './project-item';

export interface EnrollItem {
    user: UserInfo;
    project: ProjectItem;
    message: string;
}
