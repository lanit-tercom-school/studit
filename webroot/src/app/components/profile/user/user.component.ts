import { Component, OnInit } from '@angular/core';
import {ProfileService} from "../profile.service";
import {User} from "../user";

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css'],
  providers: [ProfileService]
})
export class UserComponent implements OnInit {
  user: User[];
  constructor(private userService:ProfileService) { }

  ngOnInit() {
    this.user=this.userService.getInformation();
  }

}
