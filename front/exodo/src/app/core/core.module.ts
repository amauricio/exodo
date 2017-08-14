import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LayoutComponent } from './layout/layout.component';
import { RouterModule } from "@angular/router/";
import { TopBarComponent } from './layout/top-bar/top-bar.component';
import { MainContentComponent } from './layout/main-content/main-content.component';
import { LeftBarComponent } from './layout/left-bar/left-bar.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule
  ],
  declarations: [LayoutComponent, TopBarComponent, MainContentComponent, LeftBarComponent],
  exports :[
    LayoutComponent
  ]
})
export class CoreModule { }
