export class ProjectItem {
  constructor() {
       this.Name = "";
       this.Description = "";
       this.DateOfCreation = "";
       this.Logo = "https://yegitsin.com/admin/pages/pictures/empty.jpg";
       this.Tags = "";
       this.Id = 0;
       this.Status = "";
       
    }
  Name: string;
  Description: string;
  DateOfCreation: string;
  Logo: string;
  Tags: {};
  Id: number; 
  Status: string;
}
