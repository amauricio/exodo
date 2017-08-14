import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from "@angular/router";
import { HttpModule } from "@angular/http";
import { CoreModule } from "./core/core.module";
import { LayoutComponent } from "./core/layout/layout.component";
@NgModule({
  declarations: [
  ],
  imports: [
    RouterModule,
    HttpModule,
    CoreModule,
    BrowserModule,
  ],
  providers: [],
  bootstrap: [LayoutComponent]
})
export class AppModule { }
