export class ProjectItem {
  constructor() {
    this.Name = '';
    this.Description = '';
    this.DateOfCreation = '';
    this.Logo = '';
    this.Tags = '';
    this.Id = 0;
    this.Status = '';
    this.GitHubUrl = '';
  }
  // tslint:disable-next-line:member-ordering
  Name: string;
  Description: string;
  DateOfCreation: string;
  Logo: string;
  Tags: {};
  Id: number;
  Status: string;
  GitHubUrl: string;
}
