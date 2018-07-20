import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ProjNewsComponent } from './shared-components/project-items/proj-news/proj-news.component';
import { ProjNewsItemComponent } from './shared-components/project-items/proj-news/proj-news-item/proj-news-item.component';
import { ProjectService } from 'services/project.service';
@NgModule({
    imports: [
        CommonModule
    ],
    declarations: [
        ProjNewsComponent,
        ProjNewsItemComponent
    ],
    exports: [
        ProjNewsComponent,
        ProjNewsItemComponent
    ],
    providers: [
        ProjectService
    ]
    })
    export class SharedModule {}