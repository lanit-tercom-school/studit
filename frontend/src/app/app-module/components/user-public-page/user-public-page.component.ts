import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { StudentService } from 'services/student.service';
import { UserService } from 'services/user.service';
import { CurrentUser } from 'models/current-user';
import { ProjectShort } from 'models/project-short';

@Component({
  selector: 'app-user-public-page',
  templateUrl: './user-public-page.component.html',
  styleUrls: ['./user-public-page.component.css']
})
export class UserPublicPageComponent implements OnInit {

  public CurrentUser: BehaviorSubject<CurrentUser> = new BehaviorSubject(new CurrentUser());
  public Projects: Observable<ProjectShort[]> = new Observable<ProjectShort[]>();
  public UserId: number = -1;
  constructor(
    private userService: UserService,
    private route: ActivatedRoute,
    private studentService: StudentService
  ) { }

  ngOnInit() {
    this.route.params
      .subscribe(params => {
        this.UserId = +params['id'];
        this.Projects = this.studentService.getProjectByUsers(this.UserId);
        this.userService.getUserById(this.UserId)
          .subscribe(res => {
            let c: CurrentUser = new (CurrentUser);
            c.User.Avatar = res.Avatar;
            c.User.Id = +res.Id;
            c.User.Description = res.Description;
            c.User.Nickname = res.Nickname;
            this.CurrentUser.next(c);
          });
      });
  }
}
